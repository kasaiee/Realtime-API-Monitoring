from django.db import models

class EndpointRequest(models.Model):
    name = models.CharField(max_length=30)
    method_choices = (
        (1, "GET"),
        (2, "HEAD"),
        (3, "POST"),
        (4, "PUT"),
        (5, "DELETE"),
        (6, "CONNECT"),
        (7, "OPTIONS"),
        (8, "TRACE"),
        (9, "PATCH"),
    )
    method = models.IntegerField(max_length=10,
                        choices=method_choices, default="GET")
    uri = models.CharField(max_length=100)
    header = models.CharField(max_length=500, null=True, blank=True)
    body = models.CharField(max_length=500, null=True, blank=True)
    description = models.CharField(max_length=500, null=True, blank=True)

    
class BaseURL(models.Model):
    baseUrl = models.CharField(max_length=100)

    def __str__(self):
        return self.baseUrl