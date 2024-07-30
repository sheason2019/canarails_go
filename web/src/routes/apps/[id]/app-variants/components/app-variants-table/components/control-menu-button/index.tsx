import { MouseEvent } from 'react';
import { components } from '@/api/api-gen';
import MoreHorizIcon from '@mui/icons-material/MoreHoriz';
import { IconButton, Menu, MenuItem } from '@mui/material';
import { useId, useState } from 'react';
import SetDefaultVariantMenuItem from './components/set-default-variant-menu-item';

interface Props {
  appVar: components['schemas']['AppVariant'];
}

export default function ControlMenuButton({ appVar }: Props) {
  const [anchorEl, setAnchorEl] = useState<HTMLElement | null>(null);
  const buttonId = useId();
  const menuId = useId();

  const isOpen = Boolean(anchorEl);

  const handleClick = (e: MouseEvent<HTMLElement>) => {
    setAnchorEl(e.currentTarget);
  };
  const handleClose = () => setAnchorEl(null);

  return (
    <>
      <IconButton
        aria-label="control-menu"
        aria-controls={isOpen ? menuId : undefined}
        aria-expanded={isOpen ? 'true' : undefined}
        aria-haspopup="true"
        id={buttonId}
        onClick={handleClick}
      >
        <MoreHorizIcon />
      </IconButton>
      <Menu
        id={menuId}
        MenuListProps={{ 'aria-labelledby': buttonId }}
        open={isOpen}
        anchorEl={anchorEl}
        onClose={handleClose}
      >
        <SetDefaultVariantMenuItem onClose={handleClose} appVar={appVar} />
      </Menu>
    </>
  );
}
