import AppUser from './components/app-user';
import { AppBar, Toolbar, Typography } from '@mui/material';
import StyledLink from '@/common/styled-link';
import NavigationDrawer from './components/navigation-drawer';

export default function AppHeader() {
  return (
    <AppBar position="sticky">
      <Toolbar>
        <NavigationDrawer />
        <Typography variant="h6" component="div" sx={{ flexGrow: 1, ml: 2 }}>
          <StyledLink to="/">Canarails</StyledLink>
        </Typography>
        <AppUser />
      </Toolbar>
    </AppBar>
  );
}
