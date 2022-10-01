import { Colour } from "./cube";
import { CreateFaceForColour, Face } from "./face";

export function assertSameFaceColours (expected: Face, actual: Face) {
  expect(expected.top.left.colour).toEqual(actual.top.left.colour);
  expect(expected.top.centre.colour).toEqual(actual.top.centre.colour);
  expect(expected.top.right.colour).toEqual(actual.top.right.colour);
  expect(expected.centre.left.colour).toEqual(actual.centre.left.colour);
  expect(expected.centre.centre.colour).toEqual(actual.centre.centre.colour);
  expect(expected.centre.right.colour).toEqual(actual.centre.right.colour);
  expect(expected.bottom.left.colour).toEqual(actual.bottom.left.colour);
  expect(expected.bottom.centre.colour).toEqual(actual.bottom.centre.colour);
  expect(expected.bottom.right.colour).toEqual(actual.bottom.right.colour);
}

describe("Main face rotation", () => {
  test("Clockwise face rotation", () => {
    const face = CreateFaceForColour(Colour.White);
    const expectedFace = CreateFaceForColour(Colour.White);

    face.bottom.left.colour = Colour.Red;
    face.bottom.centre.colour = Colour.Red;
    face.bottom.right.colour = Colour.Red;

    expectedFace.bottom.left.colour = Colour.Red;
    expectedFace.centre.left.colour = Colour.Red;
    expectedFace.top.left.colour = Colour.Red;

    face.Rotate(false);

    assertSameFaceColours(expectedFace, face);
  });

  test("Counter Clockwise face rotation", () => {
    const face = CreateFaceForColour(Colour.White);
    const expectedFace = CreateFaceForColour(Colour.White);

    face.bottom.left.colour = Colour.Red;
    face.bottom.centre.colour = Colour.Red;
    face.bottom.right.colour = Colour.Red;

    expectedFace.bottom.right.colour = Colour.Red;
    expectedFace.centre.right.colour = Colour.Red;
    expectedFace.top.right.colour = Colour.Red;

    face.Rotate(true);

    assertSameFaceColours(expectedFace, face);
  });
});
