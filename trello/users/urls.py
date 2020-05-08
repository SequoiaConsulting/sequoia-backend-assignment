from django.urls import path
from users.views import UserView, UserDetail, UserExist, LoginView


urlpatterns = [
    path('', UserView.as_view(), name='user_list'),
    path('<str:username>/', UserDetail.as_view(), name='user_detail'),
    path('check-existing-user/<str:username>/',
         UserExist.as_view(), name='user_exist'),
    path('auth/login/', LoginView.as_view(), name='user_login'),
]