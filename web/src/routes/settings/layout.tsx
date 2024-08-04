import { useMemo } from 'react';
import { List, Box, ListSubheader } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2/Grid2';
import { Outlet, useLocation } from '@modern-js/runtime/router';
import NavigationItem from '../components/app-header/components/navigation-item';
import KeyIcon from '@mui/icons-material/Key';
import PersonIcon from '@mui/icons-material/Person';
import InfoIcon from '@mui/icons-material/Info';

enum SettingRoutes {
  profile,
  authorization,
  about,
}

export default function Layout() {
  const location = useLocation();
  const currentRoute = useMemo(() => {
    const current = location.pathname.split('/')[2];
    switch (current) {
      case 'authorization':
        return SettingRoutes.authorization;
      case 'about':
        return SettingRoutes.about;
      default:
        return SettingRoutes.profile;
    }
  }, [location.pathname]);

  return (
    <main>
      <Grid container spacing={2}>
        <Grid md={3} xs={12}>
          <List>
            <ListSubheader>用户设置</ListSubheader>
            <NavigationItem
              icon={<PersonIcon />}
              label="用户信息"
              to="/settings"
              selected={currentRoute === SettingRoutes.profile}
            />
            <NavigationItem
              icon={<KeyIcon />}
              label="权限"
              to="/settings/authorization"
              selected={currentRoute === SettingRoutes.authorization}
            />
            <ListSubheader>应用设置</ListSubheader>
            <NavigationItem
              icon={<InfoIcon />}
              label="关于"
              to="/settings/about"
              selected={currentRoute === SettingRoutes.about}
            />
          </List>
        </Grid>
        <Grid md={9} xs={12}>
          <Box sx={{ pt: 2 }}>
            <Outlet />
          </Box>
        </Grid>
      </Grid>
    </main>
  );
}
