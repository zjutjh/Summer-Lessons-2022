<template>
  <layout>
    <k-header>
      <h1>课程表</h1>
      <k-button @click="showLoginPop">{{
        isLogin ? "当前用户:" + user.name : "登陆"
      }}</k-button>
    </k-header>
    <section class="has-sider">
      <k-sider>
        <week-selector v-model:selected-index="selectedWeek"></week-selector>
        <div @click="showAddLessonPop">添加课程</div>

        <div @click="handelFlushTable">刷新</div>
      </k-sider>
      <k-content>
        <lessons-table
          :lessons="lessons"
          :currentWeek="selectedWeek"
          @showDetailPop="showLessonDetailPop"
        ></lessons-table>
      </k-content>
    </section>
    <k-footer v-html="copyright"></k-footer>
  </layout>
  <k-modal v-model:isVisible="showPop">
    <k-card v-if="popName === 'add'" title="添加课程">
      <k-form>
        <k-input label="课程名称" type="text" v-model="lessonNew.name" />
        <k-input label="教师" type="text" v-model="lessonNew.teacher" />
        <k-input label="上课位置" type="text" v-model="lessonNew.address" />
        <k-time-selector
          label="教学周"
          :range="16"
          v-model="lessonNewTime.week"
        />
        <k-select label="周" v-model="lessonNew.week_day"></k-select>
        <k-time-selector
          label="上课时间"
          :range="12"
          v-model="lessonNewTime.time"
        />
      </k-form>
      <template #footer>
        <k-button @click="handleNewLesson">提交</k-button>
      </template>
    </k-card>

    <k-card v-if="popName === 'login'" title="登录">
      <k-form>
        <k-input
          v-if="isRegister"
          label="用户名"
          v-model="user.name"
          type="text"
        />
        <k-input label="Email" v-model="user.email" type="text" />
        <k-input
          v-if="isRegister"
          label="手机号"
          v-model="user.phone"
          type="text"
        />
        <k-input label="密码" v-model="user.password" type="password" />
        <k-input
          v-if="isRegister"
          label="验证密码"
          v-model="passwordRepeat"
          type="password"
        />
      </k-form>
      <template #footer>
        <k-button v-if="!isRegister" @click="handleLogin">登录</k-button>
        <k-button v-else @click="isRegister = false">返回</k-button>
        <k-button @click="isRegister = true" v-if="!isRegister">注册</k-button>
        <k-button v-else @click="handleRegister">提交</k-button>
      </template>
    </k-card>

    <k-card v-if="popName === 'detail'" title="课程详情">
      <k-form v-show="editLesson">
        <k-input label="课程名称" type="text" v-model="lessonDetail!.name" />
        <k-input label="教师" type="text" v-model="lessonDetail!.teacher" />
        <k-input label="上课位置" type="text" v-model="lessonDetail!.address" />
        <k-time-selector
          label="教学周"
          :range="16"
          v-model="lessonDetailTime.week"
        />
        <k-select label="周" v-model="lessonDetail!.week_day" />
        <k-time-selector
          label="上课时间"
          :range="12"
          v-model="lessonDetailTime.time"
        />
      </k-form>
      <ul v-show="!editLesson">
        <li>课程名称: {{ lessonDetail!.name }}</li>
        <li>教师: {{ lessonDetail!.teacher }}</li>
        <li>上课地点: {{ lessonDetail!.address }}</li>
      </ul>
      <template #footer>
        <k-button @click="handleEdit" v-if="!editLesson">修改</k-button>
        <k-button @click="handleDelete" v-if="!editLesson">删除</k-button>
        <k-button @click="editSubmit" v-if="editLesson">提交</k-button>
        <k-button @click="editLesson = false" v-if="editLesson">取消</k-button>
      </template>
    </k-card>
  </k-modal>
</template>

<script lang="ts">
import { ref, defineComponent, computed, watch } from "vue";
import { getLessonsTableAPI } from "@/apis/getLessonsTable";
import LessonsTable from "@/components/LessonsTable/index.vue";
import WeekSelector from "@/components/WeekSelector/index.vue";
import Layout from "@/components/layout";
import { LessonType } from "@/types/LessonsTable";
import { Button as KButton } from "@/components/button";
import { Modal as KModal } from "@/components/modal/src";
import { Card as KCard } from "@/components/card";
import { Form as KForm } from "@/components/form";
import { Input as KInput } from "@/components/form";
import "@/styles/index.scss";
import { UserType } from "./types/User";
import { login } from "./apis/login";
import { register } from "./apis/register";
import { TimeSelector as KTimeSelector } from "@/components/form";
import { Select as KSelect } from "@/components/form";
import { getUserInfo } from "./apis/getUserInfo";
import { newLesson } from "./apis/newLesson";
import { toNumber } from "@vue/shared";
import { updateLesson } from "./apis/updateLesson";
import { deleteLesson } from "./apis/deleteLesson";
const {
  Header: KHeader,
  Footer: KFooter,
  Sider: KSider,
  Content: KContent,
} = Layout;

export default defineComponent({
  name: "App",
  components: {
    LessonsTable,
    WeekSelector,
    KHeader,
    Layout,
    KFooter,
    KSider,
    KContent,
    KButton,
    KModal,
    KCard,
    KForm,
    KInput,
    KTimeSelector,
    KSelect,
  },
  setup() {
    const selectedWeek = ref(1);
    const showPop = ref(false);
    const popName = ref<"login" | "detail" | "add">();
    const lessons = ref<LessonType[]>();
    const lessonDetail = ref<LessonType>();
    const lessonDetailTime = ref({
      week: [] as number[],
      time: [] as number[],
    });
    const token = ref<string>("");
    const isRegister = ref(false);
    const user = ref<UserType>({
      name: "",
      password: "",
      email: "",
      phone: "",
      school: "",
      stu_id: "",
    });
    const passwordRepeat = ref<string>("");
    const isLogin = ref<boolean>(false);
    const lessonNew = ref<LessonType>({
      name: "",
      teacher: "",
      address: "",
      week: "",
      time: "",
      week_day: "",
      class_id: "",
    });
    const lessonNewTime = ref({
      week: [] as number[],
      time: [] as number[],
    });
    function showAddLessonPop() {
      showPop.value = true;
      popName.value = "add";
    }
    function showLoginPop() {
      if (!isLogin.value) {
        showPop.value = true;
        popName.value = "login";
      }
    }
    function showLessonDetailPop(lesson: LessonType) {
      showPop.value = true;
      popName.value = "detail";
      lessonDetail.value = lesson;
    }
    function handleLogin() {
      login(user.value).then((res) => {
        if (res.data.msg === "OK") {
          isLogin.value = true;
          token.value = res.data.data.token;
          showPop.value = false;
          getUserInfo(token.value).then((res) => {
            user.value = res.data.data.info;
          });
          alert("登录成功");
          handelFlushTable();
        } else {
          alert(res.data.msg);
        }
      });
    }
    function handleRegister() {
      register(user.value).then((res) => {
        if (res.data.msg === "OK") {
          showPop.value = false;
          alert("注册成功");
        } else {
          alert(res.data.msg);
        }
      });
    }
    function handelFlushTable() {
      getLessonsTableAPI(token.value).then((res) => {
        lessons.value = res.data.data.class;
      });
    }
    function handleNewLesson() {
      lessonNewTime.value.time.sort((a, b) => a.valueOf() - b.valueOf());
      lessonNewTime.value.week.sort((a, b) => a.valueOf() - b.valueOf());
      newLesson(
        {
          ...lessonNew.value,
          week: lessonNewTime.value.week.join(","),
          time: lessonNewTime.value.time.join(","),
        },
        token.value
      ).then((res) => {
        if (res.data.msg === "OK") {
          showPop.value = false;
          alert("添加成功");
          handelFlushTable();
        } else {
          alert(res.data.msg);
        }
      });
    }
    const editLesson = ref<boolean>(false);
    function handleEdit() {
      editLesson.value = true;
      handelFlushTable();
      lessonDetailTime.value.time = lessonDetail.value?.time
        .split(",")
        .map(toNumber)!;
      lessonDetailTime.value.week = lessonDetail.value?.week
        .split(",")
        .map(toNumber)!;
    }
    function editSubmit() {
      lessonDetailTime.value.time.sort((a, b) => a.valueOf() - b.valueOf());
      lessonDetailTime.value.week.sort((a, b) => a.valueOf() - b.valueOf());
      updateLesson(
        {
          ...lessonDetail.value!,
          week: lessonDetailTime.value.week.join(","),
          time: lessonDetailTime.value.time.join(","),
        },
        token.value
      ).then((res) => {
        if (res.data.msg === "OK") {
          showPop.value = false;
          alert("修改成功");
          handelFlushTable();
        } else {
          alert(res.data.msg);
        }
      });
    }
    function handleDelete() {
      if (confirm("确定删除?")) {
        deleteLesson(lessonDetail.value!, token.value).then((res) => {
          if (res.data.msg === "OK") {
            showPop.value = false;
            alert("删除成功");
            handelFlushTable();
          } else {
            alert(res.data.msg);
          }
        });
      }
    }
    const copyright = computed({
      get() {
        return `Copyright <a href="https://github.com/zjutjh">精弘网络技术部<a> @ ${new Date().getFullYear()}`;
      },
      set(val) {},
    });
    watch(showPop, () => {
      if (showPop.value == false) {
        editLesson.value = false;
        isRegister.value = false;
      }
    });
    return {
      lessons,
      copyright,
      selectedWeek,
      showPop,
      popName,
      lessonDetail,
      showAddLessonPop,
      showLoginPop,
      showLessonDetailPop,
      user,
      isLogin,
      passwordRepeat,
      isRegister,
      handleLogin,
      handleRegister,
      handelFlushTable,
      lessonNew,
      lessonNewTime,
      token,
      handleNewLesson,
      handleEdit,
      handleDelete,
      editLesson,
      lessonDetailTime,
      editSubmit,
    };
  },
});
</script>

<style lang="scss">
.has-sider {
  display: flex;
  flex-direction: row;
  flex: 1;
}
</style>
