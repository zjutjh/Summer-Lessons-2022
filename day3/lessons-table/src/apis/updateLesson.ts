import { LessonType } from '@/types/LessonsTable';
import { myAxios } from './axios';

export function updateLesson(lesson: LessonType, token: string) {
  return myAxios({
    url: "/api/class",
    method: "put",
    data: lesson,
    headers: {
      'Authorization': token,
    }
  });
}