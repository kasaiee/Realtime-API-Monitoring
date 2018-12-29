from django.db import models

class EndpointRequest(models.Model):
    name = models.CharField(max_length=30)
    method = models.CharField(max_length=10)
    uri = models.CharField(max_length=100)
    header = models.Charfield(max_length=500)


    
