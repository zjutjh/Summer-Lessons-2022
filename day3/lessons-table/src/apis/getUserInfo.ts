import { myAxios } from "./axios";

export function getUserInfo(token: string) {
  return myAxios({
    url: "/api/user",
    method: "get",
    headers: {
      "Authorization": token
    }
  });
}