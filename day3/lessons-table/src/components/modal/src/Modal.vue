<template>
  <div v-show="isVisible" class="k-modal">
    <div class="k-modal-mask" v-show="mask" @click="closeModal"></div>
    <div class="k-modal-content">
      <slot></slot>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
const modalProps = {
  mask: {
    type: Boolean,
    default: true,
  },
  isVisible: {
    type: Boolean,
    default: false,
  },
};
export default defineComponent({
  name: "Modal",
  props: modalProps,
  setup(props, context) {
    function closeModal() {
      context.emit("update:isVisible", false);
    }
    return {
      closeModal,
    };
  },
});
</script>

<style lang="scss">
.k-modal {
  position: fixed;
  top: 0;
  left: 0;
  z-index: 3000;
  width: 100vw;
  height: 100vh;
}
.k-modal-mask {
  background: rgb(0, 0, 0, 0.65);
  position: fixed;
  left: 0;
  top: 0;
  height: 100%;
  width: 100%;
}
.k-modal-content {
  z-index: 3100;
  position: relative;
  top: 0;
  margin: 0 auto;
  width: 400px;
  top: 30vh;
  > div {
    margin: 0 auto;
    border-radius: 4px;
  }
}
</style>
