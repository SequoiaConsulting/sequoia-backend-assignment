from django.urls import path
from tasks.views import TaskView, ParticularTaskView


urlpatterns = [
    path('', TaskView.as_view(), name='task'),
    path('<int:tid>/', ParticularTaskView.as_view(), name='particular_task'),
]