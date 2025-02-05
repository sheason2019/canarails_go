import { useId } from 'react';
import { components } from '@/api/api-gen';
import MoreHorizIcon from '@mui/icons-material/MoreHoriz';
import { IconButton, Menu } from '@mui/material';
import SetDefaultVariantMenuItem from './components/set-default-variant-menu-item';
import useDialog from '@/common/use-dialog';

interface Props {
  appVar: components['schemas']['AppVariant'];
}

export default function ControlMenuButton({ appVar }: Props) {
  const { isOpen, anchorEl, onClose, onOpen } = useDialog();

  const buttonId = useId();
  const menuId = useId();

  return (
    <>
      <IconButton
        aria-label="control-menu"
        aria-controls={isOpen ? menuId : undefined}
        aria-expanded={isOpen ? 'true' : undefined}
        aria-haspopup="true"
        id={buttonId}
        onClick={onOpen}
      >
        <MoreHorizIcon />
      </IconButton>
      <Menu
        id={menuId}
        MenuListProps={{ 'aria-labelledby': buttonId }}
        open={isOpen}
        anchorEl={anchorEl}
        onClose={onClose}
      >
        <SetDefaultVariantMenuItem onClose={onClose} appVar={appVar} />
      </Menu>
    </>
  );
}
