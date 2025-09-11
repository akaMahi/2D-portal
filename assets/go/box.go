components {
  id: "box"
  component: "/assets/scripts/box.script"
}
embedded_components {
  id: "co"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.0\n"
  "group: \"box\"\n"
  "mask: \"default\"\n"
  "mask: \"player\"\n"
  "mask: \"box\"\n"
  "mask: \"feet\"\n"
  "mask: \"head\"\n"
  "mask: \"op\"\n"
  "mask: \"bp\"\n"
  "mask: \"white_wall\"\n"
  "mask: \"black_wall\"\n"
  "mask: \"trigger\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"box\"\n"
  "  }\n"
  "  data: 16.0\n"
  "  data: 16.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "label"
  type: "label"
  data: "size {\n"
  "  x: 15.0\n"
  "  y: 16.0\n"
  "}\n"
  "color {\n"
  "  x: 0.0\n"
  "  y: 0.0\n"
  "  z: 0.0\n"
  "}\n"
  "text: \"E\"\n"
  "font: \"/builtins/fonts/default.font\"\n"
  "material: \"/builtins/fonts/label-df.material\"\n"
  ""
  position {
    y: 33.0
    z: 0.1
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"portal_box\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 32.0\n"
  "  y: 32.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/atlas/props.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.2
  }
}
