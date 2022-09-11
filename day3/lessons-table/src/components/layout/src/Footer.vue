<template>
  <footer class="k-footer">
    <slot> </slot>
  </footer>
</template>
<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  setup() {},
});
</script>
<style lang="scss" scoped>
.k-footer {
  padding: 24px 50px;
  color: #000000d9;
  font-size: 14px;
  background: #f0f2f5;
}
</style>
