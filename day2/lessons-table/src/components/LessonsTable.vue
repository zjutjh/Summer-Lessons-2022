<template>
  <div class="lessons-table">
    <div></div>
    <div class="weekday-index-panel">
      <div v-for="index in 7" :key="index">{{ index }}</div>
    </div>
    <div class="jc-index-panel">
      <div v-for="index in 12" :key="index">{{ index }}</div>
    </div>
    <div class="table">
      <template v-for="lesson in lessonsList">
        <template v-for="(time, index) in lesson.class_times" :key="index">
          <div class="lesson-card-wrapper" :style="getPosition(time)">
            <div
              class="lesson-card"
              :style="getDynamicColor(lesson.name.charCodeAt(0))"
            >
              <div class="address">{{ lesson.address }}</div>
              <div class="name">{{ lesson.name }}</div>
            </div>
          </div>
        </template>
      </template>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { LessonType, ClassTimeType } from "../types/LessonTable";

export default defineComponent({
  props: {
    lessonsList: Array,
    currentWeek: Number,
  },
  mounted() {
    console.log(this.lessonsList);
  },
  computed: {
    lessonTable(): LessonType[] | undefined {
      return this.lessonsList;
    },
  },

  methods: {
    getDynamicColor(seed: number) {
      const colorSet = ["ef8d6c", "58cc9d", "68a9df", "fad65f", "f7ac83"];
      return `
        background-color: #${colorSet[seed % colorSet.length]};`;
    },
    getPosition(timeItem: ClassTimeType) {
      const begin = parseInt(timeItem.time.split(",")[0]);
      const end = parseInt(timeItem.time.split(",").slice(-1)[0]);
      const weekday = parseInt(timeItem.week_day);
      const fontSize = Math.min(12, (end - begin + 2) * 4) + "px";
      const height = ((end - begin + 1) * 100) / 12 + "%";
      const top = "calc(" + ((begin - 1) * 100) / 12 + "%)";
      const left = "calc(" + ((weekday - 1) * 100) / 7 + "%)";
      return `top: ${top};
                left: ${left};
               height: ${height};
                font-size: ${fontSize};`;
    },
  },
});
</script>

<style scoped>
.lessons-table {
  height: 100%;
  display: grid;
  grid-template-columns: 32px 1fr;
  grid-template-rows: 32px 1fr;
}
.weekday-index-panel {
  display: flex;
  width: 100%;
  justify-content: space-around;
  border-bottom: 1px solid rgb(0, 0, 0, 0.06);
  border-left: 1px solid rgb(0, 0, 0, 0.06);
  box-sizing: border-box;
  text-align: center;
}
.jc-index-panel {
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  border-top: 1px solid rgb(0, 0, 0, 0.06);
  border-right: 1px solid rgb(0, 0, 0, 0.06);
  box-sizing: border-box;
  text-align: center;
}
.table {
  width: 100%;
  height: 100%;
  position: relative;
}
.lesson-card-wrapper {
  position: absolute;
  padding: 0.5rem 0 0 0.5rem;
  width: calc(100% / 7);
  text-align: center;
  box-sizing: border-box;
}
.lesson-card {
  height: 100%;
  justify-content: center;
}
</style>
