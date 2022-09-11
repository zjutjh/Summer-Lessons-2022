import { UserType } from '@/types/User';
import { myAxios } from './axios';

export function login(user: UserType) {
  return myAxios({
    url: "/api/user/login",
    method: "post",
    data: user,
  });
}