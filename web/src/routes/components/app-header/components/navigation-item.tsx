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
  onClose(): void;
  target?: HTMLAttributeAnchorTarget;
  suffix?: ReactNode;
}

export default function NavigationItem({
  icon,
  label,
  suffix,
  to,
  target,
  onClose,
}: Props) {
  return (
    <StyledLink to={to} target={target} onClick={onClose}>
      <ListItem disablePadding>
        <ListItemButton>
          <ListItemIcon>{icon}</ListItemIcon>
          <ListItemText primary={label} />
          {suffix}
        </ListItemButton>
      </ListItem>
    </StyledLink>
  );
}
