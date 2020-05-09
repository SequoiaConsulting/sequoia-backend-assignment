# Generated by Django 2.2.4 on 2020-05-08 12:07

from django.conf import settings
from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        migrations.swappable_dependency(settings.AUTH_USER_MODEL),
        ('projects', '0005_auto_20200508_1154'),
    ]

    operations = [
        migrations.CreateModel(
            name='Task',
            fields=[
                ('tid', models.AutoField(primary_key=True, serialize=False)),
                ('title', models.CharField(max_length=255)),
                ('description', models.TextField(blank=True, null=True)),
                ('due_date', models.DateTimeField(blank=True, null=True)),
                ('status', models.CharField(max_length=255)),
                ('creation_timestamp', models.DateTimeField(auto_now_add=True)),
                ('assignee', models.ForeignKey(blank=True, null=True, on_delete=django.db.models.deletion.SET_NULL, related_name='assignee', to=settings.AUTH_USER_MODEL, to_field='username')),
                ('assigner', models.ForeignKey(blank=True, null=True, on_delete=django.db.models.deletion.SET_NULL, related_name='assigner', to=settings.AUTH_USER_MODEL, to_field='username')),
                ('pbid', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='projects.ProjectBoard')),
            ],
        ),
    ]
