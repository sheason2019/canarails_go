import { useId } from 'react';
import MoreHorizIcon from '@mui/icons-material/MoreHoriz';
import { IconButton, Menu } from '@mui/material';
import useDialog from '@/common/use-dialog';
import { UserToken } from '../../../../typings';
import DeleteUserToken from './components/delete-user-token';

interface Props {
  userToken: UserToken;
}

export default function UserTokenItemMenuButton({ userToken }: Props) {
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
        size="small"
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
        <DeleteUserToken userToken={userToken} onClose={onClose} />
      </Menu>
    </>
  );
}
