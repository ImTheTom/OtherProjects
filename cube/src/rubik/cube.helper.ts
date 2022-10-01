import { Face } from "./face";

export function rotateFront (top: Face, bottom: Face, left: Face, right: Face, counterClockwise: boolean) {
  if (counterClockwise) {
    return rotateFrontCounterClockwiseEdges(top, bottom, left, right);
  }
  return rotateFrontClockwiseEdges(top, bottom, left, right);
}

function rotateFrontClockwiseEdges (top: Face, bottom: Face, left: Face, right: Face) {
  const tmpBot = JSON.parse(JSON.stringify(bottom));

  bottom.top.left = right.bottom.left;
  bottom.top.centre = right.centre.left;
  bottom.top.right = right.top.left;

  right.bottom.left = top.bottom.right;
  right.centre.left = top.bottom.centre;
  right.top.left = top.bottom.left;

  top.bottom.right = left.top.right;
  top.bottom.centre = left.centre.right;
  top.bottom.left = left.bottom.right;

  left.top.right = tmpBot.top.left;
  left.centre.right = tmpBot.top.centre;
  left.bottom.right = tmpBot.top.right;

  return {
    top,
    bottom,
    left,
    right
  };
}

function rotateFrontCounterClockwiseEdges (top: Face, bottom: Face, left: Face, right: Face) {
  const tmpBot = JSON.parse(JSON.stringify(bottom));

  bottom.top.left = left.top.right;
  bottom.top.centre = left.centre.right;
  bottom.top.right = left.bottom.right;

  left.top.right = top.bottom.right;
  left.centre.right = top.bottom.centre;
  left.bottom.right = top.bottom.left;

  top.bottom.right = right.bottom.left;
  top.bottom.centre = right.centre.left;
  top.bottom.left = right.top.left;

  right.top.left = tmpBot.top.right;
  right.centre.left = tmpBot.top.centre;
  right.bottom.left = tmpBot.top.left;

  return {
    top,
    bottom,
    left,
    right
  };
}

export function rotateBack (top: Face, bottom: Face, left: Face, right: Face, counterClockwise: boolean) {
  if (counterClockwise) {
    return rotateBackCounterClockwiseEdges(top, bottom, left, right);
  }
  return rotateBackClockwiseEdges(top, bottom, left, right);
};

function rotateBackClockwiseEdges (top: Face, bottom: Face, left: Face, right: Face) {
  const tmpBot = JSON.parse(JSON.stringify(bottom));

  bottom.bottom.left = left.top.left;
  bottom.bottom.centre = left.centre.left;
  bottom.bottom.right = left.bottom.left;

  left.top.left = top.top.right;
  left.centre.left = top.top.centre;
  left.bottom.left = top.top.left;

  top.top.right = right.bottom.right;
  top.top.centre = right.centre.right;
  top.top.left = right.top.right;

  right.bottom.right = tmpBot.bottom.left;
  right.centre.right = tmpBot.bottom.centre;
  right.top.right = tmpBot.bottom.right;

  return {
    top,
    bottom,
    left,
    right
  };
}

function rotateBackCounterClockwiseEdges (top: Face, bottom: Face, left: Face, right: Face) {
  const tmpBot = JSON.parse(JSON.stringify(bottom));

  bottom.bottom.right = right.top.right;
  bottom.bottom.centre = right.centre.right;
  bottom.bottom.left = right.bottom.right;

  right.top.right = top.top.left;
  right.centre.right = top.top.centre;
  right.bottom.right = top.top.right;

  top.top.left = left.bottom.left;
  top.top.centre = left.centre.left;
  top.top.right = left.top.left;

  left.bottom.left = tmpBot.bottom.right;
  left.centre.left = tmpBot.bottom.centre;
  left.top.left = tmpBot.bottom.left;

  return {
    top,
    bottom,
    left,
    right
  };
}

export function rotateLeft (top: Face, bottom: Face, front: Face, back: Face, counterClockwise: boolean) {
  if (counterClockwise) {
    return rotateLeftCounterClockwiseEdges(top, bottom, front, back);
  }
  return rotateLeftClockwiseEdges(top, bottom, front, back);
};

function rotateLeftClockwiseEdges (top: Face, bottom: Face, front: Face, back: Face) {
  const tmpBot = JSON.parse(JSON.stringify(bottom));

  bottom.top.left = front.top.left;
  bottom.centre.left = front.centre.left;
  bottom.bottom.left = front.bottom.left;

  front.top.left = top.top.left;
  front.centre.left = top.centre.left;
  front.bottom.left = top.bottom.left;

  top.top.left = back.bottom.right;
  top.centre.left = back.centre.right;
  top.bottom.left = back.top.right;

  back.top.right = tmpBot.bottom.left;
  back.centre.right = tmpBot.centre.left;
  back.bottom.right = tmpBot.top.left;

  return {
    top,
    bottom,
    front,
    back
  };
}

function rotateLeftCounterClockwiseEdges (top: Face, bottom: Face, front: Face, back: Face) {
  const tmpBot = JSON.parse(JSON.stringify(bottom));

  bottom.bottom.left = back.top.right;
  bottom.centre.left = back.centre.right;
  bottom.top.left = back.bottom.right;

  back.top.right = top.bottom.left;
  back.centre.right = top.centre.left;
  back.bottom.right = top.top.left;

  top.bottom.left = front.bottom.left;
  top.centre.left = front.centre.left;
  top.top.left = front.top.left;

  front.bottom.left = tmpBot.bottom.left;
  front.centre.left = tmpBot.centre.left;
  front.top.left = tmpBot.top.left;

  return {
    top,
    bottom,
    front,
    back
  };
}

export function rotateRight (top: Face, bottom: Face, front: Face, back: Face, counterClockwise: boolean) {
  if (counterClockwise) {
    return rotateRightCounterClockwiseEdges(top, bottom, front, back);
  }
  return rotateRightClockwiseEdges(top, bottom, front, back);
};

function rotateRightClockwiseEdges (top: Face, bottom: Face, front: Face, back: Face) {
  const tmpBot = JSON.parse(JSON.stringify(bottom));

  bottom.top.right = back.bottom.left;
  bottom.centre.right = back.centre.left;
  bottom.bottom.right = back.top.left;

  back.bottom.left = top.top.right;
  back.centre.left = top.centre.right;
  back.top.left = top.bottom.right;

  top.top.right = front.top.right;
  top.centre.right = front.centre.right;
  top.bottom.right = front.bottom.right;

  front.top.right = tmpBot.top.right;
  front.centre.right = tmpBot.centre.right;
  front.bottom.right = tmpBot.bottom.right;

  return {
    top,
    bottom,
    front,
    back
  };
}

function rotateRightCounterClockwiseEdges (top: Face, bottom: Face, front: Face, back: Face) {
  const tmpBot = JSON.parse(JSON.stringify(bottom));

  bottom.bottom.right = front.bottom.right;
  bottom.centre.right = front.centre.right;
  bottom.top.right = front.top.right;

  front.bottom.right = top.bottom.right;
  front.centre.right = top.centre.right;
  front.top.right = top.top.right;

  top.top.right = back.bottom.left;
  top.centre.right = back.centre.left;
  top.bottom.right = back.top.left;

  back.top.left = tmpBot.bottom.right;
  back.centre.left = tmpBot.centre.right;
  back.bottom.left = tmpBot.top.right;

  return {
    top,
    bottom,
    front,
    back
  };
}

export function rotateBottom (front: Face, back: Face, left: Face, right: Face, counterClockwise: boolean) {
  if (counterClockwise) {
    return rotateBottomCounterClockwiseEdges(front, back, left, right);
  }
  return rotateBottomClockwiseEdges(front, back, left, right);
};

function rotateBottomClockwiseEdges (front: Face, back: Face, left: Face, right: Face) {
  const tmpFront = JSON.parse(JSON.stringify(front));

  front.bottom.left = left.bottom.left;
  front.bottom.centre = left.bottom.centre;
  front.bottom.right = left.bottom.right;

  left.bottom.left = back.bottom.left;
  left.bottom.centre = back.bottom.centre;
  left.bottom.right = back.bottom.right;

  back.bottom.left = right.bottom.left;
  back.bottom.centre = right.bottom.centre;
  back.bottom.right = right.bottom.right;

  right.bottom.left = tmpFront.bottom.left;
  right.bottom.centre = tmpFront.bottom.centre;
  right.bottom.right = tmpFront.bottom.right;

  return {
    front,
    back,
    left,
    right
  };
}

function rotateBottomCounterClockwiseEdges (front: Face, back: Face, left: Face, right: Face) {
  const tmpFront = JSON.parse(JSON.stringify(front));

  front.bottom.left = right.bottom.left;
  front.bottom.centre = right.bottom.centre;
  front.bottom.right = right.bottom.right;

  right.bottom.left = back.bottom.left;
  right.bottom.centre = back.bottom.centre;
  right.bottom.right = back.bottom.right;

  back.bottom.left = left.bottom.left;
  back.bottom.centre = left.bottom.centre;
  back.bottom.right = left.bottom.right;

  left.bottom.left = tmpFront.bottom.left;
  left.bottom.centre = tmpFront.bottom.centre;
  left.bottom.right = tmpFront.bottom.right;

  return {
    front,
    back,
    left,
    right
  };
}

export function rotateTop (front: Face, back: Face, left: Face, right: Face, counterClockwise: boolean) {
  if (counterClockwise) {
    return rotateTopCounterClockwiseEdges(front, back, left, right);
  }
  return rotateTopClockwiseEdges(front, back, left, right);
};

function rotateTopClockwiseEdges (front: Face, back: Face, left: Face, right: Face) {
  const tmpFront = JSON.parse(JSON.stringify(front));

  front.top.left = right.top.left;
  front.top.centre = right.top.centre;
  front.top.right = right.top.right;

  right.top.left = back.top.left;
  right.top.centre = back.top.centre;
  right.top.right = back.top.right;

  back.top.left = left.top.left;
  back.top.centre = left.top.centre;
  back.top.right = left.top.right;

  left.top.left = tmpFront.top.left;
  left.top.centre = tmpFront.top.centre;
  left.top.right = tmpFront.top.right;

  return {
    front,
    back,
    left,
    right
  };
}

function rotateTopCounterClockwiseEdges (front: Face, back: Face, left: Face, right: Face) {
  const tmpFront = JSON.parse(JSON.stringify(front));

  front.top.left = left.top.left;
  front.top.centre = left.top.centre;
  front.top.right = left.top.right;

  left.top.left = back.top.left;
  left.top.centre = back.top.centre;
  left.top.right = back.top.right;

  back.top.left = right.top.left;
  back.top.centre = right.top.centre;
  back.top.right = right.top.right;

  right.top.left = tmpFront.top.left;
  right.top.centre = tmpFront.top.centre;
  right.top.right = tmpFront.top.right;

  return {
    front,
    back,
    left,
    right
  };
}
