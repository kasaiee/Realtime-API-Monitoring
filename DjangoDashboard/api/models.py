from django.db import models

class EndpointRequest(models.Model):
    name = models.CharField(max_length=30)
    method_choices = (
        ("GET", 0),
        ("HEAD", 1),
        ("POST", 2),
        ("PUT", 3),
        ("DELETE", 4),
        ("CONNECT", 5),
        ("OPTIONS", 6),
        ("TRACE", 7),
        ("PATCH", 8),
    )
    method = models.CharField(max_length=10, 
                        choices=method_choices, default="GET")
    uri = models.CharField(max_length=100)
    header = models.CharField(max_length=500)
    body = models.CharField(max_length=500)
    description = models.CharField(max_length=500)

    
