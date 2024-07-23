import useUser from '@/common/user/use-user';
import { Link } from '@modern-js/runtime/router';
import { Avatar, Button } from '@mui/material';

export default function AppUser() {
  const { isLogined } = useUser();

  if (!isLogined) {
    return (
      <Button color="inherit">
        <Link to="/login">登录</Link>
      </Button>
    );
  }

  return (
    <Link to="/login">
      <Avatar />
    </Link>
  );
}
