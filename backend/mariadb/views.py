import json
import random
import datetime

from django.core import serializers
from django.http import HttpResponse, JsonResponse
from mariadb.models import Test


DATA_VOLUME = 100
TIME_DELTA = datetime.timedelta(minutes=1)


# Create your views here.
def get_all_data(request):
    data = Test.objects.all()
    if len(data) == 0:
        generate_data()
        data = Test.objects.all()
    response = {
        'msg': 'success',
        'data': json.loads(serializers.serialize('json', data))
    }
    return JsonResponse(response)


def generate_data():
    start_time = datetime.datetime(year=2023, month=6, day=9, hour=8, minute=0, second=0)

    data = []
    for i in range(DATA_VOLUME):
        cur_time = start_time + i * TIME_DELTA
        cur_value = round(random.uniform(20.00, 30.00), 2)
        temp = Test(time=cur_time, value=cur_value)
        data.append(temp)

    Test.objects.bulk_create(data, 50)
