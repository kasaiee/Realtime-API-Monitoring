from django.db import models

class EndpointRequest(models.Model):
    name = models.CharField(max_length=30)
    url = models.CharField(max_length=100)
    
