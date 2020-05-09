from projects.models import ProjectBoard, UserProjectRelation
from rest_framework import serializers


class ProjectBoardSerializer(serializers.ModelSerializer):
    class Meta:
        model = ProjectBoard
        fields = '__all__'


class UserProjectRelationSerializer(serializers.ModelSerializer):
    class Meta:
        model = UserProjectRelation
        fields = '__all__'
