from django.db import models

class EndpointRequest(models.Model):
    name = models.CharField(max_length=30)
    method_choices = (
        ("GET", 0),
        ("HEAD", 1),
    )
    uri = models.CharField(max_length=100)
    header = models.CharField(max_length=500)
    body = models.CharField(max_length=500)
    description = models.CharField(max_length=500)

    
