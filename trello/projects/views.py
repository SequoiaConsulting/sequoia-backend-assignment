from projects.models import ProjectBoard, UserProjectRelation
from projects.serializers import (
    ProjectBoardSerializer, UserProjectRelationSerializer)
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.parsers import MultiPartParser, FormParser
from rest_framework import status
from trello.permissions import Permission
from rest_framework.parsers import JSONParser
from trello.api_exceptions import CustomApiException
from django.db import transaction


class ProjectView(APIView):
    parser_classes = (MultiPartParser, FormParser)

    PUB = []
    CON = ['POST']
    SEC = ['GET']

    permission_classes = (Permission,)

    # Create a new project board. This method will automatically make the
    # requesting user as an admin of this project board.
    def post(self, request, *args, **kwargs):
        data = JSONParser().parse(request)
        data['created_by'] = request.user.username

        with transaction.atomic():
            serializer = ProjectBoardSerializer(data=data)
            if serializer.is_valid():
                serializer.save()

                new_data = {}
                new_data['pbid'] = serializer.data['pbid']
                new_data['user'] = serializer.data['created_by']
                new_data['admin'] = True
                up_serializer = UserProjectRelationSerializer(data=new_data)
                if up_serializer.is_valid():
                    up_serializer.save()
                else:
                    raise CustomApiException(400, up_serializer.errors)
            else:
                raise CustomApiException(400, serializer.errors)

            return Response(serializer.data, status=status.HTTP_201_CREATED)

    # Get all users of all project boards. Application admin and
    # staff (ex: owner/developer of application) have permission to this.
    # Note: Application admin is different from a project board admin.
    def get(self, request):
        users = UserProjectRelation.objects.all()
        return Response(UserProjectRelationSerializer(users, many=True).data,
                        status=status.HTTP_200_OK)


class ParticularProjectView(APIView):
    parser_classes = (MultiPartParser, FormParser)

    PUB = []
    CON = ['PUT']
    SEC = []

    permission_classes = (Permission,)

    # Update a particular project board
    # (Only admins of this project board can do). Includes:
    # 1. Archive/Unarchive a board.
    # 2. Update statues list of a board
    #   (add/remove a status by sending a new statues list).
    # 3. Update title, description of a board.
    def put(self, request, pbid, format=None):
        if(UserProjectRelation.objects.filter(
                pbid=pbid, user=request.user.username, admin=True).exists() or
                request.user.staff):
            data = JSONParser().parse(request)
            pb = ProjectBoard.objects.get(pbid=pbid)
            serializer = ProjectBoardSerializer(pb, data=data, partial=True)
            if serializer.is_valid():
                serializer.save()
                return Response(
                            serializer.data, status=status.HTTP_202_ACCEPTED)
            else:
                raise CustomApiException(400, serializer.errors)
        else:
            raise CustomApiException(403, "Permission Denied")


class UserProjectView(APIView):
    parser_classes = (MultiPartParser, FormParser)

    PUB = []
    CON = ['POST', 'PUT', 'GET']
    SEC = []

    permission_classes = (Permission,)

    # Add new user to User Project Relation for a particular project board.
    # (Only admins of this project board can do)
    def post(self, request, pbid, **kwargs):
        if(UserProjectRelation.objects.filter(
                pbid=pbid, user=request.user.username, admin=True).exists() or
                request.user.staff):
            data = JSONParser().parse(request)
            data['pbid'] = pbid
            serializer = UserProjectRelationSerializer(data=data)
            if serializer.is_valid():
                serializer.save()
                return Response(
                    serializer.data, status=status.HTTP_201_CREATED)
            else:
                raise CustomApiException(400, serializer.errors)
        else:
            raise CustomApiException(403, "Permission Denied")

    # Update User Project Relation for a particular project board. Includes:
    # 1. Deactivating a user (will also make user a non-admin).
    # 2. Making Non-admin user to admin or vice-versa.
    # Note:An admin can't deactivate or can't become non-admin himself/herself
    #       Another admin of that board is required to do so.
    def put(self, request, pbid, format=None):
        if(UserProjectRelation.objects.filter(
                pbid=pbid, user=request.user.username, admin=True).exists() or
                request.user.staff):
            data = JSONParser().parse(request)
            if('user' not in data):
                raise CustomApiException(400, "Bad Request")
            elif(request.user.username == data['user']):
                raise CustomApiException(406, "Not acceptable on oneself")

            upid = UserProjectRelation.objects.get(
                    pbid=pbid, user=data['user'])

            serializer = UserProjectRelationSerializer(
                            upid, data=data, partial=True)
            if serializer.is_valid():
                serializer.save()
                if not upid.active:
                    upid.admin = False
                    upid.save()
                return Response(
                            serializer.data, status=status.HTTP_202_ACCEPTED)
            else:
                raise CustomApiException(400, serializer.errors)
        else:
            raise CustomApiException(403, "Permission Denied")

    # Retrieve all users for a particular project board with few other details
    # Admin/Non-admin user of only that project board have permission for this
    def get(self, request, pbid):
        if(UserProjectRelation.objects.filter(
                pbid=pbid, user=request.user.username).exists() or
                request.user.staff):
            users = UserProjectRelation.objects.filter(pbid=pbid)
            serializer = UserProjectRelationSerializer(users, many=True)
            return Response(serializer.data, status=status.HTTP_200_OK)
        else:
            raise CustomApiException(403, "Permission Denied")


class AllUserProjectView(APIView):
    parser_classes = (MultiPartParser, FormParser)

    PUB = []
    CON = ['GET']
    SEC = []

    permission_classes = (Permission,)

    # Retrieves all users of all the projects with which a user is associated.
    # If a user is member of n projects then he/she can see all users of those
    # n projects. (same like whatsapp group or github group)
    def get(self, request):
        user = request.user.username
        projects = UserProjectRelation.objects.filter(
                            user=user).values_list('pbid', flat=True)

        projects = list(projects)
        response = []
        for project in projects:
            users = UserProjectRelation.objects.filter(
                            pbid=project).values_list('user', flat=True)

            new_dict = {}
            new_dict['project_id'] = project
            new_dict['users'] = list(users)
            response.append(new_dict)

        return Response(response, status=status.HTTP_200_OK)
