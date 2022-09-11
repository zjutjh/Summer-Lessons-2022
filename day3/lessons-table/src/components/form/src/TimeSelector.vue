<template>
  <div class="k-time-selector">
    <div class="label">{{ label }}</div>
    <button class="btn" @click="handleBtn" type="button">
      {{ isOpen ? "收起" : "选择" }}
    </button>
    <div class="selector" v-show="isOpen" ref="selector">
      <div class="item" id="all">
        <div class="label" id="all">全选</div>
        <input
          type="checkbox"
          class="checkbox"
          @change="handleAllCheck(($event.target as HTMLInputElement).checked)"
        />
      </div>
      <br />
      <div class="item" v-for="i in range">
        <div class="label">{{ i }}</div>
        <input
          type="checkbox"
          :value="i"
          @change="handleCheck(i, ($event.target as HTMLInputElement).checked)"
          v-model="checkBoxes[i]"
        />
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.k-time-selector {
  display: grid;
  grid-template-areas: "a b" "c c";

  .label {
    width: 80px;
  }
  .selector {
    display: flex;
    flex-wrap: wrap;
    width: 90%;
    // margin: auto;
    border: 1px solid;
    flex-grow: center;
    .item {
      width: 60px;
      display: flex;
      &#all {
        width: 100%;
      }
      .label {
        width: 18px;
        height: 24px;
        &#all {
          width: 50px;
        }
      }
    }
  }
}
</style>

<script lang="ts">
import { defineComponent, onMounted, ref, VNodeRef, watch } from "vue";

export default defineComponent({
  name: "TimeSelector",
  props: {
    range: Number,
    modelValue: Array<Number>,
    label: String,
  },
  setup(props, context) {
    const checkedValue = ref<Array<Number>>([]);
    const isOpen = ref<Boolean>(false);
    const checkBoxes = ref<Array<boolean>>(new Array(props.range! + 1));
    console.log(checkBoxes);
    function updateCheck() {
      context.emit("update:modelValue", checkedValue.value);
      for (let i = 1; i <= props.range!; i++) {

        if (checkedValue.value.includes(i)) {
          checkBoxes.value[i] = true
        } else {
          checkBoxes.value[i] = false
        }
      }
    }
    function handleCheck(value: number, isCheck: boolean) {
      if (isCheck) {
        checkedValue.value.push(value);
      } else {
        checkedValue.value.splice(checkedValue.value.indexOf(value), 1);
      }
      context.emit("update:modelValue", checkedValue.value);
    }
    function handleAllCheck(isCheck: boolean) {
      if (isCheck) {
        checkedValue.value = Array.from(
          Array((props.range as number) + 1).keys()
        ).slice(1);
        updateCheck();
      } else {
        checkedValue.value = [];
        updateCheck();
      }
    }
    function handleBtn() {
      isOpen.value = !isOpen.value;
      checkedValue.value = props.modelValue!;
      updateCheck();
    }
    return {
      handleCheck,
      handleAllCheck,
      checkedValue,
      isOpen,
      handleBtn,
      // checkboxes,
      checkBoxes,
    };
  },
});
</script>
