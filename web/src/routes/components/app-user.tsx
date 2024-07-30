import StyledLink from '@/common/styled-link';
import useUser from '@/common/user/use-user';
import { Avatar, Button } from '@mui/material';

export default function AppUser() {
  const { isLogined } = useUser();

  if (!isLogined) {
    return (
      <Button color="inherit">
        <StyledLink to="/login">登录</StyledLink>
      </Button>
    );
  }

  return (
    <StyledLink to="/login">
      <Avatar />
    </StyledLink>
  );
}
