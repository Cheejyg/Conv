package com.github.cheejyg.conv.java;

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import com.github.cheejyg.conv.java.*;
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * @author Cheejyg
 */
public class Conv {
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// ToString
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * ToString parses <code>Object</code> object_ to <code>String</code>.
	 * 
	 * @author Cheejyg
	 * @param object_ The <code>Object</code> to be parsed.
	 * @return <code>String</code>
	 */
	public static String ToString(Object object_) {
		return object_.toString();
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Byte conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * ByteToShort parses <code>byte</code> byte_ to <code>Short</code>.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>Short</code>
	 */
	public static Short ByteToShort(byte byte_) {
		Short short_ = new Short(byte_);

		return short_;
	}

	/**
	 * ByteToByte returns <code>Byte</code> byte_.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>Byte</code>
	 */
	public static Byte ByteToByte(byte byte_) {
		return byte_;
	}

	/**
	 * ByteToBoolean parses <code>byte</code> byte_ to <code>Boolean</code>.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>Boolean</code>
	 */
	public static Boolean ByteToBoolean(byte byte_) {
		if (byte_ == 0) {
			return false;
		} else if (byte_ == 1) {
			return true;
		} else {
		}

		return null;
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Boolean conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * BooleanToString parses <code>boolean</code> boolean_ to <code>String</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>String</code>
	 */
	public static String BooleanToString(boolean boolean_) {
		if (boolean_ == false) {
			return "false";
		} else if (boolean_ == true) {
			return "true";
		} else {
		}

		return null;
	}

	/**
	 * BooleanToLong parses <code>boolean</code> boolean_ to <code>Long</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>Long</code>
	 */
	public static Long BooleanToLong(boolean boolean_) {
		if (boolean_ == false) {
			return 0L;
		} else if (boolean_ == true) {
			return 1L;
		} else {
		}

		return null;
	}

	/**
	 * BooleanToInteger parses <code>boolean</code> boolean_ to
	 * <code>Integer</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>Integer</code>
	 */
	public static Integer BooleanToInteger(boolean boolean_) {
		if (boolean_ == false) {
			return 0;
		} else if (boolean_ == true) {
			return 1;
		} else {
		}

		return null;
	}

	/**
	 * BooleanToShort parses <code>boolean</code> boolean_ to <code>Short</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>Short</code>
	 */
	public static Short BooleanToShort(boolean boolean_) {
		if (boolean_ == false) {
			return 0;
		} else if (boolean_ == true) {
			return 1;
		} else {
		}

		return null;
	}

	/**
	 * BooleanToByte parses <code>boolean</code> boolean_ to <code>Byte</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>Byte</code>
	 */
	public static Byte BooleanToByte(boolean boolean_) {
		if (boolean_ == false) {
			return 0;
		} else if (boolean_ == true) {
			return 1;
		} else {
		}

		return null;
	}

	/**
	 * BooleanToBoolean returns <code>Boolean</code> boolean_.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be returned.
	 * @return <code>Boolean</code>
	 */
	public static Boolean BooleanToBoolean(boolean boolean_) {
		return boolean_;
	}
}
