from tasks.models import Task
from projects.models import UserProjectRelation
from tasks.serializers import TaskSerializer
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.parsers import MultiPartParser, FormParser
from rest_framework import status
from trello.permissions import Permission
from rest_framework.parsers import JSONParser
from trello.api_exceptions import CustomApiException
from django.shortcuts import get_object_or_404


class TaskView(APIView):
    parser_classes = (MultiPartParser, FormParser)

    PUB = []
    CON = ['POST']
    SEC = []

    permission_classes = (Permission,)

    # Creates a new task for given project board. Permission to only those
    # users who are associated with this project board.
    # The status of task (case-sensitive) must belong to this project's
    # statuses. Otherwise it will raise error.
    def post(self, request, *args, **kwargs):
        data = JSONParser().parse(request)
        serializer = TaskSerializer(data=data)
        if serializer.is_valid():
            if(UserProjectRelation.objects.filter(
                    pbid=data['pbid'], user=request.user.username).exists() or
                    request.user.staff):
                serializer.save()
                return Response(serializer.data,
                                status=status.HTTP_201_CREATED)
            else:
                raise CustomApiException(403, "Permission Denied")
        raise CustomApiException(400, serializer.errors)


class ParticularTaskView(APIView):
    parser_classes = (MultiPartParser, FormParser)

    PUB = []
    CON = ['PUT', 'DELETE', 'GET']
    SEC = []

    permission_classes = (Permission,)

    # Updates/Edits a particular task. Permission to only those users
    # who are associated with the task's project board. Task's project
    # board id can't be edited.
    def put(self, request, tid, format=None):
        data = JSONParser().parse(request)
        task = get_object_or_404(Task, tid=tid)
        project = task.pbid
        if(UserProjectRelation.objects.filter(
                pbid=project.pbid, user=request.user.username).exists() or
                request.user.staff):
            data['pbid'] = project.pbid
            serializer = TaskSerializer(task, data=data, partial=True)
            if serializer.is_valid():
                serializer.save()
                return Response(serializer.data,
                                status=status.HTTP_202_ACCEPTED)
            else:
                raise CustomApiException(400, serializer.errors)
        else:
            raise CustomApiException(403, "Permission Denied")

    # Deletes a particular task. Permission to only those users
    # who are associated with the task's project board.
    def delete(self, request, tid, format=None):
        task = get_object_or_404(Task, tid=tid)
        project = task.pbid
        if(UserProjectRelation.objects.filter(
                pbid=project.pbid, user=request.user.username).exists() or
                request.user.staff):
            task.delete()
            return Response("Deleted", status=status.HTTP_204_NO_CONTENT)
        else:
            raise CustomApiException(403, "Permission Denied")

    # Retrieves details of a particular task. Permission to only those users
    # who are associated with the task's project board.
    def get(self, request, tid):
        task = get_object_or_404(Task, tid=tid)
        project = task.pbid
        if(UserProjectRelation.objects.filter(
                pbid=project.pbid, user=request.user.username).exists() or
                request.user.staff):
            return Response(TaskSerializer(task).data,
                            status=status.HTTP_200_OK)
        else:
            raise CustomApiException(403, "Permission Denied")
