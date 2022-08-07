export type LessonType = {
  name: string;
  teacher: string;
  address: string;
  class_times: ClassTimeType[];
};
export type ClassTimeType = {
  class_id: number;
  class_time_id: number;
  week: string;
  time: string;
  week_day: string;
};
