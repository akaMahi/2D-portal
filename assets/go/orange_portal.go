components {
  id: "portal"
  component: "/assets/scripts/portal.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"portal_orange\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/atlas/props.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.1
  }
  scale {
    x: 2.0
    y: 2.0
  }
}
embedded_components {
  id: "co"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"op\"\n"
  "mask: \"player\"\n"
  "mask: \"box\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  data: 30.0\n"
  "}\n"
  ""
}
