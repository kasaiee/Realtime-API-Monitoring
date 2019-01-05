from django.db import models

class EndpointRequest(models.Model):
    name = models.CharField(max_length=30)
    method_choices = (
        (0, "GET"),
        (1, "HEAD"),
        (2, "POST"),
        (3, "PUT"),
        (4, "DELETE"),
        (5, "CONNECT"),
        (6, "OPTIONS"),
        (7, "TRACE"),
        (8, "PATCH"),
    )
    method = models.CharField(max_length=10, 
                        choices=method_choices, default="GET")
    uri = models.CharField(max_length=100)
    header = models.CharField(max_length=500)
    body = models.CharField(max_length=500)
    description = models.CharField(max_length=500)

    
class BaseURL(models.Model):
    baseUrl = models.CharField(max_length=100)