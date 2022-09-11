<template>
  <div class="k-select">
    <div class="label">{{ label }}</div>
    <select
      :value="modelValue"
      @input="$emit('update:modelValue', ($event.target as HTMLSelectElement).value)"
    >
      <option v-for="(i, index) in range" :value="index + 1">{{ i }}</option>
    </select>
  </div>
</template>

<style lang="scss" scoped>
.k-select {
  display: flex;
  flex-direction: row;
  .label {
    width: fit-content;
  }
}
</style>
<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  props: {
    label: String,
    range: {
      type: Array<String>,
      default: () => ["一", "二", "三", "四", "五", "六", "日"],
    },
    modelValue: String,
  },
  setup(props, context) {
    return {};
  },
});
</script>
