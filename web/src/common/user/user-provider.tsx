import { PropsWithChildren, useEffect, useRef } from 'react';
import useToken from './use-token';

export default function UserProvider({ children }: PropsWithChildren) {
  const initRef = useRef(false);
  const { token, setToken } = useToken();

  // 初始化 token
  useEffect(() => {
    const localToken = localStorage.getItem('Authorization');
    if (localToken) {
      setToken(localToken);
    }

    initRef.current = true;
  }, []);

  // 持久化 token
  useEffect(() => {
    if (!initRef.current) return;

    localStorage.setItem('Authorization', token);
  }, [token]);

  return <>{children}</>;
}
