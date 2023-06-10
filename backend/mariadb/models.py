from django.db import models


# Create your models here.
class Test(models.Model):
    time = models.DateTimeField()
    value = models.DecimalField(max_digits=10, decimal_places=2)
