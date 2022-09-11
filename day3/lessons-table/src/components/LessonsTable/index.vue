<template>
  <div class="lessons-table">
    <div></div>
    <div class="weekday-index-panel index-panel">
      <div v-for="(col, index) in 7" class="num-index" :key="index">
        {{ col }}
      </div>
    </div>
    <div class="jc-index-panel index-panel">
      <div v-for="(row, index) in 12" class="num-index" :key="index">
        {{ row }}
      </div>
    </div>
    <div class="table" v-if="lessonsList">
      <template v-for="(lesson, index) in lessonsList">
        <div
          v-if="lesson.week.split(',').includes(currentWeek.toString())"
          class="flex class-card-wrapper"
          :style="getPosition(lesson)"
        >
          <card
            class="class-card"
            :style="getDynamicColor(lesson.name.charCodeAt(0))"
            @click="handleClick(lesson)"
            :key="index"
          >
            <div class="title">{{ lesson.address }}</div>
            <div class="item-content">{{ lesson.name }}</div>
          </card>
        </div>
      </template>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { LessonType } from "@/types/LessonsTable";
import { Card } from "@/components/card";
import "./index.scss";

const colorSet = ["ef8d6c", "58cc9d", "68a9df", "fad65f", "f7ac83"];
export default defineComponent({
  components: {
    Card,
  },
  props: {
    lessons: Array,
    currentWeek: {
      type: Number,
      default: 1,
    },
  },

  computed: {
    lessonsList(): LessonType[] | undefined {
      return this.lessons as LessonType[];
    },
  },
  methods: {
    getDynamicColor(seed: number) {
      return `
      background-color: #${colorSet[seed % colorSet.length]};`;
    },

    getPosition(timeItem: LessonType) {
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
    handleClick(lesson: LessonType) {
      this.$emit("showDetailPop", lesson);
    },
  },
});
</script>
