from django.contrib import admin
from .models import EndpointRequest, BaseURL

admin.site.register(EndpointRequest)
admin.site.register(BaseURL)