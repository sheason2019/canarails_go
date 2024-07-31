import React, { useState } from 'react';

export default function useDialog() {
  const [anchorEl, setAnchorEl] = useState<HTMLElement | null>(null);
  const isOpen = anchorEl !== null;

  const onOpen = (e?: React.MouseEvent<HTMLElement>) =>
    setAnchorEl(e?.currentTarget ?? document.body);
  const onClose = () => setAnchorEl(null);

  return {
    isOpen,
    anchorEl,
    onOpen,
    onClose,
  };
}
