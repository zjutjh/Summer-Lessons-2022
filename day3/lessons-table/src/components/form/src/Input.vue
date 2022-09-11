<template>
  <div class="k-input">
    <div class="label">{{ label }}</div>
    <input
      :type="type"
      :value="modelValue"
      @input="$emit('update:modelValue', ($event.target as HTMLInputElement).value)"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
export default defineComponent({
  name: "Input",
  props: {
    label: String,
    modelValue: String,
    type: String,
  },
  setup(props, context) {
    return {};
  },
});
</script>

<style lang="scss" scoped>
.k-input {
  // width: 90%;
  display: flex;
  flex-direction: row;
  column-gap: 5px;
  .label {
    width: 80px;
  }
  input {
    width: 80%;
  }
}
</style>
