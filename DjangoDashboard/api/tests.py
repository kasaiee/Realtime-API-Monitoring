from django.test import TestCase
from api.models import BaseURL, EndpointRequest

class BaseURLTestCase(TestCase):
    def setUp(self):
        BaseURL.objects.create(baseUrl="localhost")

    def test_base_url(self):
        baseURL = BaseURL.objects.get(baseUrl="localhost")
        self.assertEqual(baseURL, baseURL)

class EndpointRequestTestCase(TestCase):
    pass