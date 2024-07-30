import StyledLink from '@/common/styled-link';
import useUser from '@/common/user/use-user';
import { Avatar, Button } from '@mui/material';

export default function AppUser() {
  const { isLogined } = useUser();

  if (!isLogined) {
    return (
      <StyledLink to="/login">
        <Button color="inherit">登录</Button>
      </StyledLink>
    );
  }

  return (
    <StyledLink to="/login">
      <Avatar />
    </StyledLink>
  );
}
