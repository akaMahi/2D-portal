components {
  id: "door"
  component: "/assets/scripts/door.script"
  properties {
    id: "sp"
    value: "0.0, 209.0, 0.0"
    type: PROPERTY_TYPE_VECTOR3
  }
  properties {
    id: "fp"
    value: "0.0, 305.0, 0.0"
    type: PROPERTY_TYPE_VECTOR3
  }
}
embedded_components {
  id: "co"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"black_wall\"\n"
  "mask: \"default\"\n"
  "mask: \"player\"\n"
  "mask: \"box\"\n"
  "mask: \"feet\"\n"
  "mask: \"head\"\n"
  "mask: \"proj\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 16.0\n"
  "  data: 48.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"close\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 32.0\n"
  "  y: 96.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/atlas/props.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.1
  }
}
embedded_components {
  id: "s"
  type: "sound"
  data: "sound: \"/assets/SFX/door_open.wav\"\n"
  "speed: 0.6\n"
  ""
}
embedded_components {
  id: "s1"
  type: "sound"
  data: "sound: \"/assets/SFX/door_closed.wav\"\n"
  "speed: 0.6\n"
  ""
}
