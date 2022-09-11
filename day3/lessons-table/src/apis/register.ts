import { UserType } from '@/types/User';
import { myAxios } from './axios';

export function register(user: UserType) {
  return myAxios({
    url: "/api/user/register",
    method: "post",
    data: user,
  });
}
