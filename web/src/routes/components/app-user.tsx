import StyledLink from '@/common/styled-link';
import useUser from '@/common/user/use-user';
import { Avatar, Button } from '@mui/material';

export default function AppUser() {
  const { isLogined } = useUser();

  if (!isLogined) {
    return (
      <Button color="inherit" href="/login">
        登录
      </Button>
    );
  }

  return (
    <StyledLink to="/login">
      <Avatar />
    </StyledLink>
  );
}
