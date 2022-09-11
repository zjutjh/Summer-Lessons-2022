<template>
  <button class="k-button"><slot></slot></button>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  setup() {},
});
</script>
<style lang="scss" scoped>
.k-button {
  font-size: 14px;
  line-height: 1.6;
  padding: 4px 15px;
  border-radius: 2px;
  border: 1px solid #d9d9d9;

  box-sizing: border-box;
  background: #fff;
  transition: color 0.3s;
  cursor: pointer;
}
.k-button:hover {
  color: #40a9ff;
  border: 1px solid #40a9ff;
}
</style>
