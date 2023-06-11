<template>
  <n-grid cols="2 s:2 m:1 l:1 xl:1">
    <n-gi>
      <div class="logo-div">
        <a href="https://mariadb.org/" target="_blank">
          <img src="/mariadb.svg" class="logo" alt="MariaDB" />
        </a>
        <a href="https://www.djangoproject.com/" target="_blank">
          <img src="/django.svg" class="logo vue" alt="Django" />
        </a>
      </div>
    </n-gi>

    <n-gi>
      <div class="title-div">
        <h1 style="display: inline">MariaDB Connect APP</h1>
      </div>
    </n-gi>
  </n-grid>

  <div class="charts-div">
    <n-grid
      cols="4 s:4 m:2 l:1 xl:1 2xl:1"
      x-gap="12"
      y-gap="20"
      style="width: 95%"
    >
      <n-gi span="2">
        <n-card title="Pie Chart" size="huge">
          <v-chart
            autoresize
            :option="optionPie"
            style="width: 100%; height: 300px"
          />
        </n-card>
      </n-gi>
      <n-gi span="2">
        <n-card title="Scatter Chart" size="huge">
          <v-chart
            autoresize
            :option="optionScatter"
            style="width: 100%; height: 300px"
          />
        </n-card>
      </n-gi>
      <n-gi span="4">
        <n-card title="Line Chart" size="huge">
          <v-chart autoresize :option="optionLine" :style="{ width, height }" />
        </n-card>
      </n-gi>
      <n-gi span="4">
        <n-card title="Bar Chart" size="huge">
          <v-chart autoresize :option="optionBar" :style="{ width, height }" />
        </n-card>
      </n-gi>
    </n-grid>
  </div>
</template>

<script setup lang="ts">
import Axios from "axios";
import { onMounted, ref } from "vue";
import { use } from "echarts/core";
import { SVGRenderer } from "echarts/renderers";
import { PieChart, ScatterChart, LineChart, BarChart } from "echarts/charts";
import {
  LegendComponent,
  TitleComponent,
  GridComponent,
  TooltipComponent,
} from "echarts/components";
import VChart from "vue-echarts";

use([
  SVGRenderer,
  PieChart,
  ScatterChart,
  LineChart,
  BarChart,
  LegendComponent,
  TitleComponent,
  GridComponent,
  TooltipComponent,
]);

onMounted(() => {
  getData();
});

type ScatterChartData = [number, number][];
type PieChartData = { name: string; value: number }[];

const timeList = ref<any[]>();
const valueList = ref<any[]>();
const lineChartData = ref<any[]>();
const pieChartData = ref<PieChartData>();
const scatterChartData = ref<ScatterChartData>();

const resToScatterChartData = (data: any[]): ScatterChartData => {
  const res = data.map((x: any, index: number) => {
    return [index, +x.fields.value] as [number, number];
  });
  return res;
};
const resToPieChartData = (data: any[]): PieChartData => {
  const res: PieChartData = [
    {
      name: "[20-22)",
      value: 0,
    },
    {
      name: "[22-24)",
      value: 0,
    },
    {
      name: "[24-26)",
      value: 0,
    },
    {
      name: "[26-28)",
      value: 0,
    },
    {
      name: "[28-30]",
      value: 0,
    },
  ];

  data.forEach((elem: any) => {
    const x = elem.fields.value;
    if (x >= 20 && x < 22) {
      res[0].value++;
    } else if (x >= 22 && x < 24) {
      res[1].value++;
    } else if (x >= 24 && x < 26) {
      res[2].value++;
    } else if (x >= 26 && x < 28) {
      res[3].value++;
    } else {
      res[4].value++;
    }
  });

  return res;
};

const http = Axios.create({
  baseURL: "http://123.56.250.165:8000/",
});
const getData = () => {
  http
    .get("/mariadb")
    .then((res: any) => {
      console.log(res);
      timeList.value = res.data.data.map((x: any) =>
        new Date(x.fields.time).toLocaleTimeString()
      );
      valueList.value = res.data.data.map((x: any) => x.fields.value);
      lineChartData.value = res.data.data.map((x: any) => [
        x.fields.time,
        x.fields.value,
      ]);
      pieChartData.value = resToPieChartData(res.data.data);
      scatterChartData.value = resToScatterChartData(res.data.data);
      console.log(scatterChartData.value);
    })
    .catch((err) => {
      console.log(err);
    });
};

const width = "100%";
const height = "450px";
const optionLine = ref<any>({
  xAxis: {
    type: "time",
  },
  yAxis: {
    scale: true,
    type: "value",
  },
  series: [
    {
      data: lineChartData,
      name: "data",
      type: "line",
    },
  ],
  tooltip: {
    show: true,
  },
});
const optionBar = ref<any>({
  xAxis: {
    type: "category",
    data: timeList,
  },
  yAxis: {
    type: "value",
  },
  series: [
    {
      data: valueList,
      name: "data",
      type: "bar",
    },
  ],
  tooltip: {
    show: true,
  },
});
const optionPie = ref<any>({
  tooltip: {
    trigger: "item",
  },
  legend: {
    right: "3%",
    orient: "vertical",
  },
  series: [
    {
      name: "Value range",
      type: "pie",
      radius: ["50%", "70%"],
      avoidLabelOverlap: false,
      emphasis: {
        label: {
          show: true,
          fontSize: 32,
          fontWeight: "bold",
        },
      },
      label: {
        show: false,
        position: "center",
      },
      labelLine: {
        show: false,
      },
      data: pieChartData,
    },
  ],
});
const optionScatter = ref<any>({
  xAxis: {},
  yAxis: {
    scale: true,
  },
  series: [
    {
      type: "scatter",
      symbolSize: 5,
      data: scatterChartData,
    },
  ],
  tooltip: {
    show: true,
  },
});
</script>

<style scoped>
.logo-div {
  display: flex;
  align-content: center;
  justify-content: center;
}

.title-div {
  display: flex;
  align-content: center;
  justify-content: center;
}
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

.charts-div {
  display: flex;
  justify-content: center;
}
</style>
