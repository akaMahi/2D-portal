components {
  id: "energy_gun"
  component: "/assets/scripts/energy_gun.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"energy_gun\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
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
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_STATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"lava\"\n"
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
  "      y: -12.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 31.7327\n"
  "  data: 19.855562\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "fac"
  type: "factory"
  data: "prototype: \"/assets/go/energy.go\"\n"
  ""
}
