import { atom, useRecoilState } from 'recoil';

const tokenState = atom({
  key: 'token-string',
  default: localStorage.getItem('Authorization') ?? '',
});

export default function useToken() {
  const [token, setToken] = useRecoilState(tokenState);

  return {
    token,
    setToken,
  };
}
