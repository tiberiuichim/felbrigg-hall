varying mediump vec4 position;
varying mediump vec2 var_texcoord0;

uniform lowp sampler2D texture_sampler;
uniform lowp vec4 tint;

void main()
{
	// Pre-multiply alpha since all runtime textures already are
	lowp vec4 tint_pm = vec4(tint.xyz * tint.w, tint.w);

	lowp vec4 texColor = texture2D(texture_sampler, var_texcoord0.xy);
	texColor.a *= 0.5;
	
	gl_FragColor = texColor * tint_pm;
}
