<template>
  <div class="layout">
    <lessons-table :lessons-list="lessonsTable"></lessons-table>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import LessonsTable from "./components/LessonsTable.vue";
import { data } from "./data";

export default defineComponent({
  components: {
    LessonsTable,
  },
  data() {
    return {
      lessonsTable: data,
    };
  },
});
</script>

<style scoped>
.layout {
  height: 100vh;
}
</style>
