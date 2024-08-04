import { PropsWithChildren, useEffect, useRef } from 'react';
import useToken from './use-token';
import useUser from './use-user';

export default function UserProvider({ children }: PropsWithChildren) {
  const { token } = useToken();
  const { mutate } = useUser();

  // 持久化 token
  useEffect(() => {
    localStorage.setItem('Authorization', token);
    mutate();
  }, [token]);

  return <>{children}</>;
}
