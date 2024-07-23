import useUser from '@/common/user/use-user';
import { Link } from '@modern-js/runtime/router';
import { Avatar, Box, Button } from '@mui/material';

export default function AppUser() {
  const { isLogined } = useUser();

  if (!isLogined) {
    return (
      <Button color="inherit">
        <Box
          component={Link}
          to="/login"
          sx={{ color: 'inherit', textDecoration: 'none' }}
        >
          登录
        </Box>
      </Button>
    );
  }

  return (
    <Link to="/login">
      <Avatar />
    </Link>
  );
}
