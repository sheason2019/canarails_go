import { Box, Drawer, IconButton, List, ListSubheader } from '@mui/material';
import MenuIcon from '@mui/icons-material/Menu';
import useDialog from '@/common/use-dialog';
import AppsIcon from '@mui/icons-material/Apps';
import BarChartIcon from '@mui/icons-material/BarChart';
import SettingsIcon from '@mui/icons-material/Settings';
import HelpOutlineIcon from '@mui/icons-material/HelpOutline';
import OpenInNewIcon from '@mui/icons-material/OpenInNew';
import NavigationItem from './navigation-item';

export default function NavigationDrawer() {
  const { isOpen, onClose, onOpen } = useDialog();

  return (
    <>
      <IconButton
        size="large"
        edge="start"
        color="inherit"
        aria-label="menu"
        onClick={onOpen}
      >
        <MenuIcon />
      </IconButton>
      <Drawer open={isOpen} onClose={onClose}>
        <Box sx={{ width: 280 }}>
          <List subheader={<ListSubheader>菜单</ListSubheader>}>
            <NavigationItem
              icon={<AppsIcon />}
              label="Apps"
              to="/"
              onClose={onClose}
            />
            <NavigationItem
              icon={<BarChartIcon />}
              label="统计"
              to="/dashboard"
              onClose={onClose}
            />
            <NavigationItem
              icon={<SettingsIcon />}
              label="设置"
              to="/settings"
              onClose={onClose}
            />
            <NavigationItem
              icon={<HelpOutlineIcon />}
              label="帮助文档"
              to="https://canarails.sheason.site"
              target="_blank"
              suffix={<OpenInNewIcon sx={{ color: 'GrayText' }} />}
              onClose={onClose}
            />
          </List>
        </Box>
      </Drawer>
    </>
  );
}
