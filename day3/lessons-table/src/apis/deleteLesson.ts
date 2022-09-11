import { LessonType } from '@/types/LessonsTable';
import { myAxios } from './axios';

export function deleteLesson(lesson: LessonType, token: string) {
  return myAxios({
    url: "/api/class?id=" + lesson.class_id,
    method: "delete",
    headers: {
      'Authorization': token,
    }
  });
}