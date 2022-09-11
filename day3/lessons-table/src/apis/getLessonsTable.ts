import { myAxios } from "./axios";

export function getLessonsTableAPI(token: string) {
  return myAxios({
    url: "/api/class",
    method: "get",
    headers: {
      "Authorization": token
    }
  });
}
