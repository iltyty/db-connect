<template>
  <div>
    <a href="https://mariadb.org/" target="_blank">
      <img src="/mariadb.svg" class="logo" alt="MariaDB" />
    </a>
    <a href="https://www.djangoproject.com/" target="_blank">
      <img src="/django.svg" class="logo vue" alt="Django" />
    </a>
  </div>

  <h1>MariaDB Connect APP</h1>

  <v-chart autoresize :option="option" :style="{ width, height }" />
</template>

<script setup lang="ts">
import Axios from "axios";
import { onMounted, ref } from "vue";
import { use } from "echarts/core";
import { SVGRenderer } from "echarts/renderers";
import { LineChart } from "echarts/charts";
import {
  TitleComponent,
  GridComponent,
  TooltipComponent,
} from "echarts/components";
import VChart from "vue-echarts";

use([SVGRenderer, LineChart, TitleComponent, GridComponent, TooltipComponent]);

onMounted(() => {
  getData();
});

const timeList = ref<[any]>();
const valueList = ref<[any]>();
const data = ref<[any]>();
const http = Axios.create({
  baseURL: "http://123.56.250.165:8000/",
});
const getData = () => {
  http
    .get("/mariadb")
    .then((res: any) => {
      console.log(res);
      timeList.value = res.data.data.map((x: any) => x.fields.time);
      valueList.value = res.data.data.map((x: any) => x.fields.value);
      data.value = res.data.data.map((x: any) => [
        x.fields.time,
        x.fields.value,
      ]);
    })
    .catch((err) => {
      console.log(err);
    });
};

const width = "600px";
const height = "450px";
const option = ref<any>({
  title: {
    left: "center",
    text: "Data Presentation",
  },
  xAxis: {
    type: "time",
  },
  yAxis: {
    scale: true,
    type: "value",
  },
  series: [
    {
      data,
      name: "data",
      type: "line",
    },
  ],
  tooltip: {
    show: true,
  },
});
</script>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
