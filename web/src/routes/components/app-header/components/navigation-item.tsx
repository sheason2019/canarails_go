import StyledLink from '@/common/styled-link';
import {
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
} from '@mui/material';
import { HTMLAttributeAnchorTarget, ReactNode } from 'react';

interface Props {
  icon: ReactNode;
  label: ReactNode;
  to: string;
  onClose?(): void;
  target?: HTMLAttributeAnchorTarget;
  suffix?: ReactNode;
  selected?: boolean;
}

export default function NavigationItem({
  icon,
  label,
  suffix,
  to,
  target,
  selected,
  onClose,
}: Props) {
  return (
    <StyledLink to={to} target={target} onClick={onClose}>
      <ListItem disablePadding>
        <ListItemButton selected={selected}>
          <ListItemIcon>{icon}</ListItemIcon>
          <ListItemText primary={label} />
          {suffix}
        </ListItemButton>
      </ListItem>
    </StyledLink>
  );
}
