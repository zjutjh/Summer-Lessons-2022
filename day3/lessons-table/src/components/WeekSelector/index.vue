<template>
  <div class="week-selector">
    <h3>周数</h3>
    <div class="week-selector-content">
      <div
        v-for="i in 20"
        :key="i"
        class="index-wrapper"
        :class="{ 'index-selected': selectedIndex === i }"
        @click="toggleWeek(i)"
      >
        <span>
          {{ i }}
        </span>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  props: {
    selectedIndex: Number,
  },
  setup(props, context) {
    function toggleWeek(week: number) {
      context.emit("update:selectedIndex", week);
    }
    return {
      toggleWeek,
    };
  },
});
</script>

<style lang="scss" scoped>
.week-selector {
  cursor: unset;
}
.week-selector-content {
  display: grid;
  grid-template-columns: repeat(4, 50px);
  grid-template-rows: repeat(5, 50px);
  .index-wrapper {
    padding: 8px;
    border-radius: 1000px;
    text-align: center;
    transition: all 0.1s;
    cursor: pointer;
    span {
      -webkit-user-select: none;
      user-select: none;
    }
  }
  .index-wrapper:hover {
    background: rgb(0, 0, 0, 0.06);
  }
  .index-selected {
    background: rgb(0, 0, 0, 0.06);
  }
}
</style>
