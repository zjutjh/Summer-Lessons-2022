import { LessonType } from '@/types/LessonsTable';
import { myAxios } from './axios';

export function newLesson(lesson: LessonType, token: string) {
  return myAxios({
    url: "/api/class",
    method: "post",
    data: {
      ...lesson,
      class_id: null,
    },
    headers: {
      'Authorization': token,
    }
  });
}